
// EXTERNAL DEPENDENCIES
import express from 'express';
import winston from 'winston';
import path from 'path';
import rp from 'request-promise';

// INTERNAL DEPENDENCIES
import Registry from '../app';
import Image from '../database/models/image.model';
import Rating from '../database/models/rating.model';

export default class Router {

    private router: express.Router;

    constructor(registry: Registry) {

        /**
         * Create express router.
         */
        this.router = express.Router();

        /**
         * @Routing
         */

        // POST /core/images
        this.router.post('/', async (req, res) => {

            console.log('New image recieved.');

            const trx = await registry.db.transaction();

            try {

                if (!req.files) {
                    throw { message: 'No file uploaded' };
                }

                const image = await Image.create({probability: -1}, { transaction: trx });
                const filename = image.id + '.jpg';
                const uploadedFile: any = req.files.image;
                const filepath = path.join("/data/images", filename);

                uploadedFile.mv(filepath, function (err: any) {
                    if (err)
                        return res.status(500).send(err);
                });

                console.log('Image successful saved.');

                const options = {
                    method: 'POST',
                    uri: `http://${process.env.AI}:${process.env.AI_PORT}/model/predict`,
                    body: {
                        id: image.id
                    },
                    json: true
                };

                await rp.post(options);
                console.log('successful sent');

                await trx.commit();
                res.sendStatus(200);

            } catch (err) {

                await trx.rollback();
                console.log(err);
                res.status(400).send(err);

            }

        })

        // POST /core/images/:id/rating
        this.router.post('/:id/rating', async (req, res) => {

            console.log('New rating recieved.');

            const trx = await registry.db.transaction();

            try {

                const id = req.params.id;
                if (!id.match(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$/i)) {
                    throw { message: 'invalid id' };
                }

                let isCrackedBool = true;
                if (!req.body.iscracked) {
                    isCrackedBool = false;
                }

                const image = await Image.findOne({ where: { id } });
                if (!image) {
                    throw { message: 'Image not found' };
                }
                await Rating.create({ imageId: image.id, isCracked: isCrackedBool }, { transaction: trx });
                console.log('Rating registered.');

                const options = {
                    method: 'POST',
                    uri: `http://${process.env.AI}:${process.env.AI_PORT}/model/train`,
                    body: {
                        id: image.id,
                        iscracked: req.body.iscracked,
                    },
                    json: true
                };

                await rp.post(options);
                console.log('Rating sent to AI');

                if(isCrackedBool) {
                    const options_msg = {
                        method: 'POST',
                        uri: `http://${process.env.MATTERMOST}:${process.env.MATTERMOST_PORT}/image`,
                        body: {
                            id: req.params.id,
                            probability: 100,
                            channel: "alerts",
                        },
                        json: true
                    };

                    await rp.post(options_msg);
                    console.log('Message sent to frontend by human');
                }

                await trx.commit();
                res.sendStatus(200);

            } catch (err) {

                await trx.rollback();
                console.log(err);
                res.status(400).send(err);

            }

        });

        // POST /core/images/:id/probability
        this.router.post('/:id/probability', async(req, res) => {

            console.log('New Probability recieved.');

            const trx = await registry.db.transaction();

            try {

                const probability = req.body.probability;
                await Image.update({ probability }, { where: { id: req.params.id } });

                if(probability <= Number(process.env.AUTOPROB)) {
                    if(probability >= Number(process.env.AUTOPROB_LOW)) {
                        const options = {
                            method: 'POST',
                            uri: `http://${process.env.MATTERMOST}:${process.env.MATTERMOST_PORT}/image`,
                            body: {
                                id: req.params.id,
                                probability: probability,
                                channel: "detection",
                            },
                            json: true
                        };
                        
                        await rp.post(options);
                        console.log('Probability sent:'+ probability);
                    } else {
                        console.log('Probability below cutoff:'+ probability);
                    }
                } else {
                    const options = {
                        method: 'POST',
                        uri: `http://${process.env.MATTERMOST}:${process.env.MATTERMOST_PORT}/image`,
                        body: {
                            id: req.params.id,
                            probability: probability,
                            channel: "alerts",
                        },
                        json: true
                    };

                    await rp.post(options);
                    console.log('Message sent to frontend by ai:'+ probability);
                }

                await trx.commit();
                res.sendStatus(200);

            } catch (err) {

                await trx.rollback();
                console.log(err);
                res.status(400).send(err);

            }

        })
    }

    public init(): express.Router {
        return this.router;
    }

}

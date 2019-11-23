
// EXTERNAL DEPENDENCIES
import express from 'express';
import winston from 'winston';
import uuidv4 from 'uuid/v4';
import fs from 'fs';
import path from 'path';


// INTERNAL DEPENDENCIES
import Registry from '../app';
import Image from '../database/models/Image.model';
import Rating from '../database/models/Rating.model';
import { UploadedFile } from 'express-fileupload';

export default class Router {

    private logger: winston.Logger;
    private router: express.Router;

    constructor(registry: Registry) {

        /**
         * Initialize module-based logger.
         */
        this.logger = registry.loggingSystem.getLogger('IMAGE');

        /**
         * Create express router.
         */
        this.router = express.Router();

        /**
         * @Routing
         */

        // POST /core/images
        this.router.post('/', async (req, res) => {

            this.logger.debug('New image recieved.');

            const trx = await registry.db.transaction();

            try {

                if (!req.files) {
                    throw { message: 'No file uploaded' };
                }

                const image = await Image.create({}, { transaction: trx });
                const filename = image.id + '.jpg';
                const uploadedFile: any = req.files.image;
                const filepath = path.join(process.env.static!, filename);

                uploadedFile.mv(filepath, function (err: any) {
                    if (err)
                        return res.status(500).send(err);
                });

                await trx.commit();
                this.logger.info('Image successful saved.');
                res.sendStatus(200);

            } catch (err) {

                await trx.rollback();
                this.logger.error(err);
                res.status(400).send(err);

            }

        })

        // POST /core/image/:id/rating
        this.router.post('/:id/rating', async (req, res) => {

            this.logger.debug('New rating recieved.');

            const trx = await registry.db.transaction();

            try {

                const id = req.params.id;
                if (!id.match(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$/i)) {
                    throw { message: 'invalid id' };
                }

                let isCracked = true;
                if (!req.body.isCracked) {
                    isCracked = false;
                }

                const image = await Image.findOne({ where: { id } });
                if (!image) {
                    throw { message: 'Image not found' };
                }
                await Rating.create({ imageId: image.id, isCracked }, { transaction: trx });

                await trx.commit();
                this.logger.info('Rating successful registered.');
                res.sendStatus(200);

            } catch (err) {

                await trx.rollback();
                this.logger.error(err);
                res.status(400).send(err);

            }

        });
    }

    public init(): express.Router {
        return this.router;
    }

}


async function createFile(filepath: string, image: any) {
    return new Promise((resolve, reject) => {
        fs.writeFile(filepath, image, (err) => {
            if (err) throw err;
            console.log('The file has been saved!');
        });
    });
}
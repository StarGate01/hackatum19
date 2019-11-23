
// EXTERNAL DEPENDENCIES
import { RequestHandler } from 'express';

export default class Responder {

    public static send(): RequestHandler {
        return (req, res) => {

            if (res.locals.response && res.locals.response.status) {

                const status = res.locals.response.status;
                const funct = res.locals.response.function;
                const message = res.locals.response.message;
                const data = res.locals.response.data;


                switch (funct) {
                    case 'sendData':
                        res.status(status).send(data);
                        break;
                    case 'sendStatus':
                        res.sendStatus(status);
                        break;
                    case 'sendFile':
                        res.status(status).sendFile(data);
                        break;
                    default:
                        res.status(status).send(message);
                        break;
                }

            } else {

                

            }

        }
    }

}
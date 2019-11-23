
// EXTERNAL DEPENDENCIES
import express from 'express';
import winston from 'winston';

// INTERNAL DEPENDENCIES
import Registry from '../app';
import { Responder } from '../middleware';

export default class Router {

    private logger: winston.Logger;
    private router: express.Router;

    constructor(registry: Registry) {

        /**
         * Initialize module-based logger.
         */
        this.logger = registry.loggingSystem.getLogger('TEST');

        /**
         * Create express router.
         */
        this.router = express.Router();

        /**
         * @Routing
         */

        // GET /api/test/
        this.router.get('/',
            (_, res, next) => {
                this.logger.debug('This is a test route.')

                res.locals.response = {
                    function: 'sendData',
                    status: 200,
                    data: {
                        comment: 'This is a test route.',
                    }
                }

                next();
            },
            Responder.send());

    }

    public init(): express.Router {
        return this.router;
    }

}
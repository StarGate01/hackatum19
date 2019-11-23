
// EXTERNAL DEPENDENCIES
import express from 'express';

// INTERNAL DEPENDENCIES
import Registry from '../app';
import TestRouter from './test-router';

export default class Router {

    private mainRouter: express.Router;

    constructor(registry: Registry) {

        /**
         * Create express router.
         */
        this.mainRouter = express.Router();

        /**
         * Set baseUrl.
         */
        registry.app.use('/core', this.mainRouter);

        /**
         * Bind router modules.
         */
        this.mainRouter.use('/test', new TestRouter(registry).init());

    }

}
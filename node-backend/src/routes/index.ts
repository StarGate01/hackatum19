
// EXTERNAL DEPENDENCIES
import express from 'express';

// INTERNAL DEPENDENCIES
import Registry from '../app';
import ImageRouter from './image-router';

export default class Router {

    private mainRouter: express.Router;

    constructor(registry: Registry) {

        /**
         * Create express router.
         */
        this.mainRouter = express.Router();

        /**
         * Bind router modules.
         */
        registry.app.use(this.mainRouter);
        this.mainRouter.use('/images', new ImageRouter(registry).init());

        /**
         * Every other undefined route returns 'Page not found'.
         */
        registry.app.use((_, res, next) => res.status(404).send({message: 'Page not found'}));

    }

}

// EXTERNAL DEPENDENCIES
import bodyParser from 'body-parser';
import dotenv from 'dotenv';
import express from 'express';
import winston from 'winston';
import fileUpload from 'express-fileupload';

// INTERNAL DEPENDENCIES
import RoutingSystem from './routes';
import Database from './database';

export default class Server {

    public app: express.Express;
    public routingSystem: RoutingSystem;
    public db: Database;

    constructor() {

        /**
         * Loads environment variables from .env file into process.env.
         */
        dotenv.config();

        /**
         * Creates an Express application.
         */
        this.app = express();

        /**
         * Parse incoming request bodies in a middleware before the
         * handlers, available under the req.body property.
         */
        this.app.use(bodyParser.json());

        /**
         * When you upload a file, the file will be accessible from req.files.
         */
        this.app.use(fileUpload({
            limits: { fileSize: 50 * 1024 * 1024 },
        }));

        /**
         * Initialize database.
         */
        this.db = new Database(this);

        /**
         * Binds routing system.
         */
        this.routingSystem = new RoutingSystem(this);

    }

    /**
     * Starts the backend with all services.
     */
    public async bootstrap(): Promise<void> {

        /**
         * Set database connection.
         */
        await this.db.initConnection();

        /**
         * Application starts listening on port.
         */
        this.app.listen(process.env.PORT, () => {
            console.log(`App listening on port ${process.env.PORT}!`);
        });

    }

}

new Server().bootstrap();
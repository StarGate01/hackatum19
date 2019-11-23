
// EXTERNAL DEPENDENCIES
import dotenv from 'dotenv';
import express from 'express';
import morgan from 'morgan';
import winston from 'winston';

// INTERNAL DEPENDENCIES
import LoggingSystem from './utils/logging-system';
import RoutingSystem from './routes';
import Database from './database';

export default class Server {

    public app: express.Express;
    public loggingSystem: LoggingSystem;
    public routingSystem: RoutingSystem;
    public db: Database;

    private logger: winston.Logger;

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
         * Binds logging system.
         */
        this.loggingSystem = new LoggingSystem('logs');

        /**
         * HTTP request logger middleware for node.js.
         */
        this.app.use(morgan(':status :remote-addr - :remote-user '
            + '":method :url HTTP/:http-version" '
            + ':status :res[content-length] '
            + '":referrer" ":user-agent"',
            { stream: this.loggingSystem.getMorganStream() })
        );

        /**
         * Initialize module-based logger.
         */
        this.logger = this.loggingSystem.getLogger('APP');

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
            this.logger.info(`App listening on port ${process.env.PORT}!`);
        });

    }

}

new Server().bootstrap();
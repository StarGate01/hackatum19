// EXTERNAL DEPENDENCIES
import path from 'path';
import { Sequelize } from 'sequelize-typescript';
import { Transaction } from 'sequelize/types';
import winston from 'winston';

// INTERNAL DEPENDENCIES
import Registry from '../app';

export default class Database {

    /**
     * @private
     *
     * @property {Object} db
     *  Postgres database instance.
     * @property {Object} dbConfig
     *  Database configuration.
     * @property {Object} logger
     *  Module-based winston logger.
     */
    private db: Sequelize;
    private dbConfig: any;
    private logger: winston.Logger;

    /**
     * @constructor
     *
     * @param {Object} registry
     *  The registry handles everything central.
     */
    constructor(registry: Registry) {

        /**
         * Create module-base winston logger.
         */
        this.logger = registry.loggingSystem.getLogger('DB');

        /**
         * Get database configuration, depending on environment.
         */
        this.dbConfig = {
            username : 'postgres',
            password: 'test123#',
            database: 'testDB',
            host: 'localhost',
            port: '5432',
            dialect: 'postgres',
            schema: 'main',
        }

        /**
         * Initialize postgres database instance.
         */
        this.db = new Sequelize({
            username: this.dbConfig.username,
            password: this.dbConfig.password,
            database: this.dbConfig.database,
            host: this.dbConfig.host,
            port: this.dbConfig.port,
            dialect: this.dbConfig.dialect,
            logging: false,
            define: {
                schema: this.dbConfig.schema,
            }
        });

    }

    /**
     * Start connecting with database.
     */
    public async initConnection(): Promise<Sequelize> {

        try {

            this.logger.info(`Initialize database configuration:`);
            this.logger.info(`Host:     ${this.dbConfig.host}`);
            this.logger.info(`Port:     ${this.dbConfig.port}`);
            this.logger.info(`Username: ${this.dbConfig.username}`);
            this.logger.info(`Database: ${this.dbConfig.database}`);
            this.logger.info(`Schema:   ${this.dbConfig.schema}`);

            /**
             * Create schema, if not existing yet.
             */
            const schemata = await this.db.showAllSchemas({});
            if (schemata.indexOf(this.dbConfig.schema) <= -1) {
                await this.db.createSchema(this.dbConfig.schema as string, {});
                await this.db.sync({});
            }

            /**
             * Add all models to database.
             */
            await this.db.addModels([path.join(__dirname, 'models', '**/*.model.js')]);
            await this.db.sync({});

            this.logger.info(`Database initialized\n`);
            return this.db;

        } catch (err) {
            this.logger.error('Failed to initialize database.');
            throw (err);
        }

    }

    /**
     * Start database transaction.
     */
    public async transaction(): Promise<Transaction> {
        return this.db.transaction();
    }
}
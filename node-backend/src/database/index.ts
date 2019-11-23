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

    /**
     * @constructor
     *
     * @param {Object} registry
     *  The registry handles everything central.
     */
    constructor(registry: Registry) {

        /**
         * Get database configuration, depending on environment.
         */
        this.dbConfig = {
            username : process.env.DB_USER,
            password: process.env.DB_PASS,
            database: process.env.DB_DB,
            host: process.env.DB,
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

            console.log(`Initialize database configuration:`);
            console.log(`Host:     ${this.dbConfig.host}`);
            console.log(`Port:     ${this.dbConfig.port}`);
            console.log(`Username: ${this.dbConfig.username}`);
            console.log(`Database: ${this.dbConfig.database}`);
            console.log(`Schema:   ${this.dbConfig.schema}`);

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

            console.log(`Database initialized\n`);
            return this.db;

        } catch (err) {
            console.log('Failed to initialize database.');
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
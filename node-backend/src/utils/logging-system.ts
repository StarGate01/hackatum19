
// EXTERNAL DEPENDENCIES
import winston from 'winston';

export default class LoggingSystem {

    /**
     * Winston Logging Levels:
     *
     *  error: 0,
     *  warn: 1,
     *  info: 2,
     *  verbose: 3,
     *  debug: 4,
     *  silly: 5,
     */

    private foldername: string;

    /**
     * @constructor
     * 
     * @param {string} foldername
     *  Foldername where logfiles should be saved.
     */
    constructor(foldername: string) {

        this.foldername = foldername;

    }

    /**
     * Creates winston module-based logger.
     * 
     * @param label 
     *  Modulename
     */
    public getLogger(label: string): winston.Logger {

        const logger = winston.createLogger({

            transports: [
                new winston.transports.Console({
                    format: winston.format.combine(
                        winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
                        winston.format.label({ label }),
                        winston.format.colorize(),
                        winston.format.printf(({ timestamp, level, label, message }) => `${timestamp} ${level}: [${label}] ${message}`),
                    ),
                    level: 'debug',
                }),
                new winston.transports.File({
                    format: winston.format.combine(
                        winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
                        winston.format.label({ label }),
                        winston.format.printf(({ timestamp, level, label, message }) => `${timestamp} ${level}: [${label}] ${message}`),
                    ),
                    level: 'info',
                    filename: `${this.foldername}/app.log`,
                }),
                new winston.transports.File({
                    format: winston.format.combine(
                        winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
                        winston.format.label({ label }),
                        winston.format.printf(({ timestamp, level, label, message }) => `${timestamp} ${level}: [${label}] ${message}`),
                    ),
                    level: 'error',
                    filename: `${this.foldername}/errors.log`,
                }),
            ]
        });

        return logger;

    }

    /**
     * Create a stream object with a 'write' function that will be used 
     * by 'morgan'.
     */
    public getMorganStream(): { write: (message: string) => void; } {

        const logger = winston.createLogger({
            transports: [
                new winston.transports.Console({
                    format: winston.format.combine(
                        winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
                        winston.format.colorize(),
                        winston.format.printf(({ timestamp, level, message }) => `${timestamp} ${level}: ${message}`),
                    ),
                    level: 'info',
                }),
                new winston.transports.File({
                    format: winston.format.combine(
                        winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
                        winston.format.printf(({ timestamp, level, message }) => `${timestamp} ${level}: ${message}`),
                    ),
                    level: 'info',
                    filename: `${this.foldername}/http.log`,
                }),
                new winston.transports.File({
                    format: winston.format.combine(
                        winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
                        winston.format.printf(({ timestamp, level, message }) => `${timestamp} ${level}: ${message}`),
                    ),
                    level: 'error',
                    filename: `${this.foldername}/errors.log`,
                })
            ],
        });

        return {
            write: (message) => {

                // get status of message (first 3 chars)
                const status = parseInt(message.substring(0, 3), 10);

                // delete first 4 chars -> XXX_
                // and last 2 chars -> \n
                message = message.substring(4, message.length - 2);

                if (status < 400) {
                    logger.info(message);
                } else {
                    logger.error(message);
                }

            }
        };

    }

}
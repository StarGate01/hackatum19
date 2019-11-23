// EXTERNAL DEPENDENCIES
import { Table, Model, Column, DataType, HasMany } from 'sequelize-typescript';

// INTERNAL DEPENDENCIES
import Rating from './rating.model';

@Table({
    timestamps: true,
    freezeTableName: true,
})
export default class Image extends Model<Image> {

    @Column({
        primaryKey: true,
        type: DataType.UUID,
        allowNull: false,
        defaultValue: DataType.UUIDV4,
        comment: "unique id for an image in uuid/v4",
    })
    id!: string;

    @Column({
        type: DataType.INTEGER,
        allowNull: false,
    })
    probability!: number;

    // ########## ########## ########## ########## ##########
    //                      Associations
    // ########## ########## ########## ########## ##########

    @HasMany(() => Rating)
    ratings!: Rating[];

}
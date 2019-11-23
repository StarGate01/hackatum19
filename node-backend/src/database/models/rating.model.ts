// EXTERNAL DEPENDENCIES
import { Table, Model, Column, DataType, ForeignKey, BelongsTo } from 'sequelize-typescript';

// INTERNAL DEPENDENCIES
import Image from './image.model';

@Table({
    timestamps: true,
    freezeTableName: true,
})
export default class Rating extends Model<Rating> {

    @Column({
        primaryKey: true,
        type: DataType.INTEGER,
        allowNull: false,
    })
    id!: number;

    @ForeignKey(() => Image)
    @Column({
        type: DataType.UUID,
    })
    imageId!: string;

    @Column({
        type: DataType.FLOAT,
    })
    value!: number;

    // ########## ########## ########## ########## ##########
    //                      Associations
    // ########## ########## ########## ########## ##########

    @BelongsTo(() => Image)
    image!: Image;

}
// EXTERNAL DEPENDENCIES
import { Table, Model, Column, DataType, ForeignKey, BelongsTo } from 'sequelize-typescript';

// INTERNAL DEPENDENCIES
import Image from './Image.model';

@Table({
    timestamps: true,
    freezeTableName: true,
})
export default class Rating extends Model<Rating> {

    @Column({
        primaryKey: true,
        type: DataType.INTEGER,
        allowNull: false,
        autoIncrement: true,
    })
    id!: number;

    @ForeignKey(() => Image)
    @Column({
        type: DataType.UUID,
    })
    imageId!: string;

    @Column({
        type: DataType.BOOLEAN,
        allowNull: false,
    })
    isCracked!: boolean;

    // ########## ########## ########## ########## ##########
    //                      Associations
    // ########## ########## ########## ########## ##########

    @BelongsTo(() => Image)
    image!: Image;

}
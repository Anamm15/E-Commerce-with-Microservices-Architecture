import db from "../config/database.js";
import { DataTypes } from "sequelize";

const Orders = db.define('order', {
    dest_address: {
        type: DataTypes.STRING,
        allowNull: false
    },
    payment_method: {
        type: DataTypes.STRING,
        allowNull: false
    },
    total_price: {
        type: DataTypes.FLOAT
    }
}, {
    freezeTableName: true,
    timestamps: true,
    paranoid: true
})

export default Orders;
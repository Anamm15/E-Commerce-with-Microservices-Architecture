import db from "../config/database.js";
import { DataTypes } from "sequelize";

const Detail_Order = db.define('detail_order', {
    id: {
        type: DataTypes.INTEGER,
        primaryKey: true,
        autoIncrement: true
    },
    qty: {
        type: DataTypes.INTEGER,
        allowNull: false
    },
    price: {
        type: DataTypes.FLOAT,
        allowNull: false
    },
    order_date: {
        type: DataTypes.DATE,
        defaultValue: DataTypes.NOW
    },
    deliveried_date: {
        type: DataTypes.DATE
    },
    order_status: {
        type: DataTypes.STRING(30)
    }
}, {
    freezeTableName: true,
    timestamps: true,
    paranoid: true
});

export default Detail_Order;
import { Sequelize } from "sequelize";
import { Orders } from "../models/index.js"


export const makeOrder = async(data, t) => {
    const order = await Orders.create(data, { transaction: t });
    return order;
}

export const createDetailOrder = async(data, t) => {
    const detailOrder = await Detail_Order.create({
        orderId: data.orderId, 
        productId: data.productId, 
        qty: data.qty, 
        price: data.price,
        storeId: data.storeId,
        status_order: process.env.STATUS_DIPESAN
    }, {transaction: t});

    return detailOrder;
}

export const updateOrder = async(orderId, t) => {
    return await Orders.update({ 
        total_price: Sequelize.literal(`(SELECT SUM(price) FROM detail_order do WHERE do.orderId = ${orderId})`)}, { 
        where: { id: orderId },
        type: Sequelize.QueryTypes.UPDATE,
        transaction: t
    });
}
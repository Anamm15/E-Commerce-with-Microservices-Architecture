import { Orders } from "../models/index.js";



export const getOrdersByStoreId = async(storeId) => {
    return await Orders.findAll({
        where: {
            storeId
        }
    });
}

export const addOrder = async(data) => {
    return await Orders.create(data);
}
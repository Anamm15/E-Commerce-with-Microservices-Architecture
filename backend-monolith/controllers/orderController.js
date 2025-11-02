import db from "../config/database"
import { makeOrder as makeOrderFromService, updateOrder as updateOrderFromService } from "../services/order";
import { getProduct, updateStockProduct } from "../services/productService";


export const makeOrder = async(req, res) => {
    const t = await db.transaction();

    try {
        const { userId, productId, qty, price, payment_method, dest_address } = req.body;
        const orderData = {
            userId,
            dest_address,
            payment_method,
        }


        const product = await getProduct(productId, t);
        if (!product) return res.status(404).json({message: "Produk tidak ditemukan"});

        const order = await makeOrderFromService(orderData, t);
        const detailOrder = await createDetailOrder({ orderId: order.id, productId, qty, price, storeId: product.storeId }, t);
        await updateOrderFromService(order.id, t);
        await updateStockProduct(product, qty, t);

        t.commit();
        res.status(201).json({message: "Data berhasil dibuat"});
    } catch (error) {
        t.rollback();
        res.status(500).json({message: "Internal server error"});
    }
}
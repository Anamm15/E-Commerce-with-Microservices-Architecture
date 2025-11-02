import { getProducts as getProductsFromService, getProduct as getProductFromService, addProduct as addProductFromService } from "../services/productService.js"
import { convertImagePathToLink } from "../utils/utils.js";


export const getProducts = async (req, res) => {
    try {
        const products = await getProductsFromService();

        if (!products || products.length === 0) {
            return res.sendStatus(404); 
        }

        products.map((product) => {
            product.photo = convertImagePathToLink(req.protocol, req.get('host'), product.photo);
        })

        res.json(products); 
    } catch (error) {
        console.error(error);
        res.status(500).json({ message: "Internal server error" });
    }
};


export const getProduct = async(req, res) => {
    try {
        const id = req.params.id;
        const product = await getProductFromService(id);
        if (!product) return res.sendStatus(404);
        product.photo = convertImagePathToLink(req.protocol, req.get('host'), product.photo);

        res.json(product);
    } catch (error) {
        res.status(500).json({ message: "Internal server error"});
    }
}

export const addProduct = async(req, res) => {
    try {
        const { name, description, price, stock, categoryId, storeId } = req.body;
        
        const imagePath = req.file ? `/uploads/products/img/${req.file.filename}` : null;

        const data = { name, description, price, stock, categoryId, storeId, photo: imagePath };
        
        await addProductFromService(data);
        res.json({ message: "Data berhasil dibuat"});
    } catch (error) {
        res.status(500).json({ message: "Internal server error"});
    }
}
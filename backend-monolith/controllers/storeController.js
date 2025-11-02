import { getStoreByUserId as getStoreByUserIdFromService, addStore as addStoreFromService } from "../services/storeService.js";


export const getStoreByUserId = async(req, res) => {
    try {
        const userId = req.params.id;
        const store = await getStoreByUserIdFromService(userId);
        if (!store) return res.sendStatus(404);
        res.json(store);
    } catch (error) {
        res.status(500).json({ message: "Internal server error"});
    }
}

export const addStore = async(req, res) => {
    try {
        const { name, description, domicile, address, phone_num, email, userId } = req.body;
        const imagePath = req.file ? `/uploads/products/img/${req.file.filename}` : null;
        const data = { name, description, domicile, address, phone_num, email, userId, photo: imagePath };

        const store = await addStoreFromService(data);
        if (!store) res.sendStatus(404);
        
        res.status(201).json({ message: "Data berhasil ditambahkan"});
    } catch (error) {
        res.status(500).json({ message: "Internal server error"});
    }
}
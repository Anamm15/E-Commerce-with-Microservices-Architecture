import { register } from "../services/userService.js"

export const registerUser = async (req, res, next) => {
    try {
        console.log(req.body);
        await register(req.body);
        res.status(201).json({ message: "Data berhasil dibuat" });
    } catch (error) {
        if (error.name === 'SequelizeUniqueConstraintError') {
            res.status(409).json({ message: 'Email sudah digunakan. Silakan gunakan email lain.' });
        } else if (error.name === 'ValidationError') {
            res.status(400).json({ message: 'Data tidak valid', details: error.errors });
        } else {
            res.status(500).json({ message: 'Gagal membuat akun', error: error.message });
        }
    }
};
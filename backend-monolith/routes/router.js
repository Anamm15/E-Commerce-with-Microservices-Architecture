import express from 'express';
import { registerUser } from '../controllers/userController.js';
import { login, logout } from '../controllers/authController.js';
import { authenticateJWT } from '../middlewares/authMiddleware.js';
import { getCategories } from '../controllers/categoryController.js';
import { addStore, getStoreByUserId } from '../controllers/storeController.js';
import upload from '../middlewares/multerMiddleware.js';
import { addProduct, getProduct, getProducts } from '../controllers/productController.js';
import { addComment, getCommentPerProduct } from '../controllers/commentController.js';
const router = express.Router();

router.get('/', (req, res) => {
    res.send('Hello World!');
});

router.post('/auth/register', registerUser);
router.post('/auth/login', login);
router.get('/auth/logout', logout);
router.get('/test', authenticateJWT, (req, res) => {
    res.json({ message: "Selamat berhasil menerapkan" });
});

router.get('/category/get', getCategories);


router.get('/store/getStoreByUserId/:id', getStoreByUserId);
router.post('/store/add', upload.single('photo'), addStore);


router.post('/product/add', upload.single('photo'), addProduct);
router.get('/product/getAllProducts', getProducts);
router.get('/product/get/:id', getProduct);


router.get('/comment/getCommentPerProduct/:id', getCommentPerProduct);
router.post('/comment/add', addComment);

export default router;
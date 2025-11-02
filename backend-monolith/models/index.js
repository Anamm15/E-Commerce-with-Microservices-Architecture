import Categories from "./category.js";
import Comments from "./comment.js";
import Detail_Order from "./detail_order.js";
import Orders from "./orders.js";
import Products from "./product.js";
import Stores from "./store.js";
import Users from "./user.js";

Users.hasOne(Stores);
Stores.belongsTo(Users);

Users.hasMany(Orders);
Orders.belongsTo(Users);

Users.hasMany(Comments);
Comments.belongsTo(Users);

Products.hasMany(Comments);
Comments.belongsTo(Products);

Orders.belongsToMany(Products, { through: Detail_Order });
Products.belongsToMany(Orders, { through: Detail_Order });

Products.belongsTo(Categories);
Categories.hasMany(Products);

Stores.hasMany(Products);
Products.belongsTo(Stores);

Detail_Order.belongsTo(Orders);
Orders.hasMany(Detail_Order);

Detail_Order.belongsTo(Products);
Products.hasMany(Detail_Order);

Detail_Order.belongsTo(Stores);
Stores.hasMany(Detail_Order);

Comments.belongsTo(Comments, { as: 'parent', foreignKey: 'parent_comment_id' });
Comments.hasMany(Comments, { as: 'replies', foreignKey: 'parent_comment_id' });


export {
    Categories,
    Comments,
    Detail_Order,
    Orders,
    Products,
    Stores,
    Users
};
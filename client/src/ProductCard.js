import React from "react";
import { FaShoppingCart } from "react-icons/fa";
import "./ProductCard.css";

const ProductCard = ({ product }) => {
    return (
        <div className="product-card">
            <picture>
                <source media="(min-width: 1024px)" srcSet={product.image.desktop} />
                <source media="(min-width: 768px)" srcSet={product.image.tablet} />
                <source media="(max-width: 767px)" srcSet={product.image.mobile} />
                <img src={product.image.thumbnail} alt={product.name} className="product-image" />
            </picture>

            <button className="add-to-cart-btn">
                <FaShoppingCart style={{ marginRight: "8px" }} />
                Add to Cart
            </button>

            <div className="product-meta">
                <div className="product-category">{product.category}</div>
                <div className="product-name">{product.name}</div>
                <div className="product-price">${product.price.toFixed(2)}</div>
            </div>
        </div>
    );
};

export default ProductCard;

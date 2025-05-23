import React, { useState } from "react";
import "../styles/ProductCard.css";
import { ReactComponent as CartIcon } from '../cart.svg';

const ProductCard = ({ product, onAddToCart }) => {
    const [quantity, setQuantity] = useState(0);

    const handleAdd = () => {
        setQuantity(1);
        onAddToCart(product);
    };

    const increment = () => {
        setQuantity(prev => prev + 1);
        onAddToCart(product);
    };

    const decrement = () => {
        setQuantity(prev => {
            const newQty = prev - 1;
            if (newQty <= 0) return 0;
            return newQty;
        });
    };

    return (
        <div className="product-card">
            <div className="image-container">
                <img className="product-image" src={product.image.desktop} alt={product.name} />

                {quantity === 0 ? (
                    <button className="add-to-cart-btn" onClick={handleAdd}>
                        <CartIcon className="cart-icon" /> Add to Cart
                    </button>
                ) : (
                    <div className="quantity-controls">
                        <button onClick={decrement}>−</button>
                        <span>{quantity}</span>
                        <button onClick={increment}>+</button>
                    </div>
                )}
            </div>

            <div className="details">
                <p className="category">{product.category}</p>
                <h2 className="title">{product.name}</h2>
                <p className="price">${product.price.toFixed(2)}</p>
            </div>
        </div>
    );
};

export default ProductCard;

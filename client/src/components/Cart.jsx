import React from "react";
import "../styles/Cart.css";

const Cart = ({ cartItems = [], onRemove }) => {
    const getTotalItems = () => cartItems.reduce((sum, item) => sum + item.quantity, 0);
    const getTotalPrice = () =>
        cartItems.reduce((sum, item) => sum + item.quantity * item.price, 0).toFixed(2);

    return (
        <div className="cart-box">
            <h2 className="cart-title">Your Cart ({getTotalItems()})</h2>
            {cartItems.map(item => (
                <div key={item.id} className="cart-entry">
                    <div className="entry-text">
                        <div className="entry-name">{item.name}</div>
                        <div className="entry-pricing">
                            <span className="entry-qty">{item.quantity}x</span>
                            <span className="entry-unit">@${item.price.toFixed(2)}</span>
                            <span className="entry-total">= ${(item.price * item.quantity).toFixed(2)}
                            </span>
                        </div>
                    </div>
                    <button className="remove-btn" onClick={() => onRemove(item.id)}>✕</button>
                </div>
            ))}

            <div className="cart-total">
                <span>Order Total</span>
                <span className="total-price">${getTotalPrice()}</span>
            </div>

            <div className="carbon-message">
                <span>🌳</span> This is a <strong>carbon-neutral</strong> delivery
            </div>

            <button className="confirm-btn">Confirm Order</button>
        </div>
    );
};

export default Cart;

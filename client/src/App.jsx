import React, { useEffect, useState } from "react";
import "./index.css";
import ProductCard from "./components/ProductCard";
import Cart from "./components/Cart";

function App() {
    const [products, setProducts] = useState([]);
    const [cartItems, setCartItems] = useState([]);

    const handleAddToCart = (product) => {
        setCartItems(prevItems => {
            const existingItem = prevItems.find(item => item.id === product.id);

            if (existingItem) {
                // increase quantity
                return prevItems.map(item =>
                    item.id === product.id
                        ? { ...item, quantity: item.quantity + 1 }
                        : item
                );
            } else {
                // add new item
                return [...prevItems, { ...product, quantity: 1 }];
            }
        });
    };

    const onRemove = (idToRemove) => {
        setCartItems((prevCart) =>
            prevCart
                .map((item) =>
                    item.id === idToRemove
                        ? { ...item, quantity: item.quantity - 1 }
                        : item
                )
                .filter((item) => item.quantity > 0)
        );
    };

    useEffect(() => {
        fetch(
            "http://localhost:8080/api/product",
            {
                method: 'GET',
                headers: {
                    "Content-Type": "application/json",
                    "x-api-key": "apitest"
                }
            }
        )
            .then(res => res.json())
            .then(data => setProducts(data))
            .catch(err => console.error("Error fetching products:", err));
    }, []);

    return (
        <div className="main-container">
            <h1 className="page-title">Desserts</h1>
            <div className="content-layout">
                <div className="product-grid">
                    {products.map((product) => (
                        <ProductCard key={product.id} product={product} onAddToCart={handleAddToCart} />
                    ))}
                </div>
                <Cart cartItems={cartItems} onRemove={onRemove} />
            </div>
        </div>
    );
}

export default App;

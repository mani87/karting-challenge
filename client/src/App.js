import React, { useEffect, useState } from "react";
import "./App.css";
import ProductGrid from "./ProductGrid";

function App() {
    const [products, setProducts] = useState([]);
    const [cart, setCart] = useState([]);

    const addToCart = (item) => {
        setCart((prev) => [...prev, item]);
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
        <div>
            <h1>Desserts</h1>
            <ProductGrid products={products} />
        </div>
    );
}

export default App;

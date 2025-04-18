import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import {Route, Routes, BrowserRouter} from 'react-router-dom';
import Home from "./Home";


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);

export default function App() {
    return(
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />}></Route>
            </Routes>
        </BrowserRouter>
    );
}
reportWebVitals();

import React from 'react';
import '../styles.css';

const Banner = () => {
    return (
        <header className="banner">
            <div className="logo">
                <img src="https://cdn.nexmonyx.com/logos/small-logo.png" alt="NexMonyx Logo" />
            </div>
            <div className="banner-text">
                <h1>
                    <span className="nex-color">Nex</span>
                    <span className="monyx-color">Monyx</span>
                </h1>
            </div>
        </header>
    );
};

export default Banner;

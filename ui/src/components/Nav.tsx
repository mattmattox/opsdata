import React from 'react';
import {Link} from "react-router-dom";

const Nav = (props: { name: string, setName: (name: string) => void }) => {
    const logout = async () => {
        var url = `http://opsdata-api:8000/api/logout`;
        await fetch(url, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        });

        props.setName('');
    }

    let menu;

    if (props.name === '') {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login" className="nav-link">Login</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/register" className="nav-link">Register</Link>
                </li>
            </ul>
        )
    } else {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login" className="nav-link" onClick={logout}>Logout</Link>
                </li>
            </ul>
        )
    }

    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <Link to="/" className="navbar-brand">Home</Link>

                <div>
                    {menu}
                </div>
            </div>
        </nav>
    );
};

export default Nav;
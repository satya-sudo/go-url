import { Link, useNavigate } from "react-router-dom";
import { Logout, CheckLogin } from "./checkLogin";
import logo from "../assets/logo.png";

export default function Navbar() {
    const navigate = useNavigate();
    const token = CheckLogin();

    const handleLogout = () => {
        Logout();
        navigate("/login");
    };

    return (
        <nav className="navbar">
            <div className="navbar-left">
                <img src={logo} alt="go-url logo" className="navbar-logo" />
                <span className="navbar-title">go-url</span>
            </div>
            <div className="navbar-right">
                <Link to="/">Dashboard</Link>
                {token ? (
                    <button className="logout-btn" onClick={handleLogout}>
                        Logout
                    </button>
                ) : (
                    <>
                        <Link to="/login">Login</Link>
                        <Link to="/signup">Signup</Link>
                    </>
                )}
            </div>
        </nav>
    );
}

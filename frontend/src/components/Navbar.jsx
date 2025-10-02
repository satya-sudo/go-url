import { Link, useNavigate } from "react-router-dom";
import {Logout, CheckLogin} from "./checkLogin";

export default function Navbar() {
    const navigate = useNavigate();
    const token = CheckLogin();

    const handleLogout = () => {
        Logout()
        navigate("/login");
    };

    return (
        <nav className="navbar">
            <Link to="/">Dashboard</Link>
            {token ? (
                <>
                    <button onClick={handleLogout}>Logout</button>
                </>
            ) : (
                <>
                    <Link to="/login">Login</Link>
                    <Link to="/signup">Signup</Link>
                </>
            )}
        </nav>
    );
}

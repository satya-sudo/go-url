import { useState, useEffect } from "react";
import api from "../api/axios";
import Modal from "../components/Modal";

export default function Dashboard() {
    const [urls, setUrls] = useState([]);
    const [error, setError] = useState("");
    const [newUrl, setNewUrl] = useState("");
    const [showModal, setShowModal] = useState(false);
    const [selectedShortCode, setSelectedShortCode] = useState(null);

    useEffect(() => {
        fetchUrls();
    }, []);

    const fetchUrls = async () => {
        try {
            const res = await api.get("/shorten/list/all");
            setUrls(res.data);
        } catch (err) {
            setError("Failed to load URLs.");
            console.error(err);
        }
    };

    const handleCreate = async (e) => {
        e.preventDefault();
        setError("");
        try {
            await api.post("/shorten/", { longUrl: newUrl });
            setNewUrl("");
            fetchUrls();
        } catch (err) {
            setError("Failed to create URL.");
            console.error(err);
        }
    };

    const handleDeleteClick = (shortCode) => {
        setSelectedShortCode(shortCode);
        setShowModal(true);
    };

    const confirmDelete = async () => {
        try {
            await api.delete(`/shorten/${selectedShortCode}`);
            setUrls(urls.filter((u) => u.shortCode !== selectedShortCode));
        } catch (err) {
            setError("Failed to delete URL.");
            console.error(err);
        }
        setShowModal(false);
        setSelectedShortCode(null);
    };

    return (
        <div className="dashboard">
            <h2 className="dashboard-title">Your Shortened URLs</h2>

            <form onSubmit={handleCreate} className="dashboard-form">
                <input
                    className="dashboard-input"
                    placeholder="Enter long URL"
                    value={newUrl}
                    onChange={(e) => setNewUrl(e.target.value)}
                />
                <button type="submit" className="dashboard-button">Shorten</button>
            </form>

            {error && <p className="dashboard-error">{error}</p>}

            {urls.length === 0 ? (
                <p className="dashboard-empty">No URLs found.</p>
            ) : (
                <div className="dashboard-table-container">
                    <table className="dashboard-table">
                        <thead>
                        <tr>
                            <th>Short Code</th>
                            <th>Long URL</th>
                            <th>Created</th>
                            <th>Expiration</th>
                            <th>Hits</th>
                            <th>Actions</th>
                        </tr>
                        </thead>
                        <tbody>
                        {urls.map((u) => (
                            <tr key={u.shortCode}>
                                <td>
                                    <a
                                        href={`http://localhost:8080/${u.shortCode}`}
                                        target="_blank"
                                        rel="noreferrer"
                                    >
                                        {u.shortCode}
                                    </a>
                                </td>
                                <td className="long-url">{u.longUrl}</td>
                                <td>{new Date(u.createdAt).toLocaleString()}</td>
                                <td>
                                    {u.expirationAt
                                        ? new Date(u.expirationAt).toLocaleString()
                                        : "-"}
                                </td>
                                <td>{u.hits}</td>
                                <td>
                                    <button
                                        className="delete-button"
                                        onClick={() => handleDeleteClick(u.shortCode)}
                                    >
                                        Delete
                                    </button>
                                </td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </div>
            )}

            <Modal
                isOpen={showModal}
                onClose={() => setShowModal(false)}
                onConfirm={confirmDelete}
                shortCode={selectedShortCode}
            />
        </div>
    );
}

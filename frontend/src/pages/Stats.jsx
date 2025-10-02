import { useState } from "react";
import api from "../api/axios";

export default function Stats() {
    const [id, setId] = useState("");
    const [stats, setStats] = useState(null);

    const fetchStats = async (e) => {
        e.preventDefault();
        const res = await api.get(`/shorten/${id}/stats`);
        setStats(res.data);
    };

    return (
        <div>
            <form onSubmit={fetchStats} className="form">
                <h2>Get Stats</h2>
                <input placeholder="Short URL ID" value={id} onChange={(e) => setId(e.target.value)} />
                <button type="submit">Fetch</button>
            </form>

            {stats && (
                <div className="stats">
                    <h3>Stats</h3>
                    <pre>{JSON.stringify(stats, null, 2)}</pre>
                </div>
            )}
        </div>
    );
}

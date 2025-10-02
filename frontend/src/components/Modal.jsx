import "./Modal.css";

export default function Modal({ isOpen, onClose, onConfirm, shortCode }) {
    if (!isOpen) return null;

    return (
        <div className="modal-overlay">
            <div className="modal">
                <h3>Delete Short URL</h3>
                <p>Are you sure you want to delete <b>{shortCode}</b>?</p>
                <div className="modal-actions">
                    <button onClick={onConfirm} className="danger">Delete</button>
                    <button onClick={onClose}>Cancel</button>
                </div>
            </div>
        </div>
    );
}
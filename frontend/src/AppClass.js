import React, {Fragment, useState} from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function AppClass() {
    const [url, setUrl] = useState('');
    const [loading, setLoading] = useState(false);
    const [result, setResult] = useState(null);
    const [error, setError] = useState('');
    const API_BASE_URL = process.env.REACT_APP_BACKEND_API_URL;

    const handleAnalyze = async () => {
        if (!url) {
            setError('Please enter a URL.');
            return;
        }
        setError('');
        setLoading(true);
        setResult(null);
        try {
            const response = await fetch(`${API_BASE_URL}/api/analyze`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ url }),
            });
            if (!response.ok) {
                const err = await response.json();
                throw new Error(err.message || 'Unknown error');
            }
            const data = await response.json();
            setResult(data);
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    return (
        <Fragment>
            <div className="container min-vh-100 d-flex align-items-center justify-content-center">
                <div className="w-100" style={{maxWidth: 420}}>
                    {error && <div className="alert alert-danger mt-3 text-center">{error}</div>}
                    <div className="bg-white p-4 p-md-5 rounded-4 shadow text-center">
                        <h1 className="mb-3 text-primary fw-bold" style={{letterSpacing: 1}}>Web Page Analizer</h1>
                        <p className="mb-4 text-secondary">Enter the URL of the web page you want to analyze</p>
                        <input
                            type="text"
                            className="form-control form-control-lg mb-3"
                            placeholder="Enter URL"
                            value={url}
                            onChange={e => setUrl(e.target.value)}
                            disabled={loading}
                        />
                        <button
                            className="btn btn-primary btn-lg w-100 fw-semibold"
                            type="button"
                            onClick={handleAnalyze}
                            disabled={loading}
                        >
                            {loading ? 'Analyzing...' : 'Analyze'}
                        </button>
                    </div>
                </div>
            </div>
           
            {result && (
                <div className="mt-4 mx-auto" style={{maxWidth: 420}}>
                    <div className="card card-body">
                        <p><b>HTML Version:</b> {result.htmlVersion}</p>
                        <p><b>Title:</b> {result.title}</p>
                        <p><b>Headings:</b> {result.headingsSummary}</p>
                        <p><b>Links:</b> {result.linksSummary}</p>
                        <p><b>Login Form:</b> {result.hasLoginForm ? 'Yes' : 'No'}</p>
                    </div>
                </div>
            )}
        </Fragment>
    )
}

export default AppClass;
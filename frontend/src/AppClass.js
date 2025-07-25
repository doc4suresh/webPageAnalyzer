import React, { Fragment, useState } from 'react';

function AppClass() {
    const [url, setUrl] = useState('');
    const [loading, setLoading] = useState(false);
    const [result, setResult] = useState(null);
    const [error, setError] = useState('');
    const API_BASE_URL = process.env.REACT_APP_BACKEND_API_URL;

    /* Backend API Call */
    const handleAnalyze = async () => {
        if (!url) {
            setError('Please enter a URL.');
            return;
        }
        setError('');
        setLoading(true);
        setResult(null);
        try {
            const response = await fetch(`${API_BASE_URL}/analyze?url=${encodeURIComponent(url)}`, {
                method: 'GET',
                headers: { 'Content-Type': 'application/json' },
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
            {/* Header */}
            <nav className="navbar navbar-expand-lg navbar-dark bg-primary shadow-sm mb-4">
                <div className="container">
                    <span className="navbar-brand d-flex align-items-center gap-2">
                        <i className="bi bi-globe2 fs-3"></i>
                        Web Page Analyzer With GO Lang
                    </span>
                </div>
            </nav>

            {/* Analyer Card */}
            <div className="container d-flex flex-column align-items-center justify-content-center" style={{ minHeight: '80vh' }}>
                {/* Input Card */}
                <div className="w-100" style={{ maxWidth: 500 }}>
                    <div className="card border-0 shadow-sm mb-4">
                        <div className="card-body p-4">
                            <h2 className="mb-3 text-center text-primary fw-bold" style={{ letterSpacing: 1 }}>
                                Analyze Any Web Page
                            </h2>
                            <p className="mb-4 text-secondary text-center">
                                Enter the URL of the web page you want to analyze.
                            </p>
                            <div className="input-group input-group-lg mb-3">
                                <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Enter URL"
                                    value={url}
                                    onChange={e => setUrl(e.target.value)}
                                    disabled={loading}
                                    autoFocus
                                />
                            </div>
                            <button
                                className="btn btn-primary btn-lg w-100 fw-semibold"
                                type="button"
                                onClick={handleAnalyze}
                                disabled={loading}
                            >
                                {loading ? <span><span className="spinner-border spinner-border-sm me-2"></span>Analyzing...</span> : 'Analyze'}
                            </button>
                            {error && <div className="alert alert-danger mt-3 text-center">{error}</div>}
                        </div>
                    </div>
                </div>

                {/* Result Card */}
                {result && (
                    <div className="mt-2 w-100" style={{ maxWidth: 850 }}>
                        <div className="card shadow border-0">
                            <div className="card-body p-4">
                                <h4 className="mb-4 text-primary fw-bold d-flex align-items-center">
                                    <i className="bi bi-file-earmark-text me-2"></i>
                                    Web Page Analysis Result
                                </h4>
                                <div className="table-responsive">
                                    <table className="table table-borderless align-middle mb-0">
                                        <tbody>
                                            <tr>
                                                <th className="text-secondary">URL</th>
                                                <td><a href={result.url} target="_blank" rel="noopener noreferrer">{result.url}</a></td>
                                            </tr>
                                            <tr>
                                                <th className="text-secondary w-25">Title</th>
                                                <td>{result.title || <span className="text-muted">No Title Found</span>}</td>
                                            </tr>
                                            <tr>
                                                <th className="text-secondary">HTML Version</th>
                                                <td>{result.HTMLVersion || <span className="text-muted">No HTML Version Found</span>}</td>
                                            </tr>
                                            <tr>
                                                <th className="text-secondary">Login Form</th>
                                                <td>{result.IsLoginForm ? <span className="badge bg-success">Yes</span> : <span className="text-muted">No Login Form Found</span>}</td>
                                            </tr>
                                            <tr>
                                                <th className="text-secondary">Headings</th>
                                                <td>
                                                    {result.headCount && Object.keys(result.headCount).length > 0 ? (
                                                        <div className="d-flex flex-wrap gap-2">
                                                            {Object.entries(result.headCount).map(([tag, count]) => (
                                                                <span className="badge bg-warning text-dark" key={tag}>{tag.toUpperCase()}: {count}</span>
                                                            ))}
                                                        </div>
                                                    ) : (
                                                        <span className="text-muted">None Headings Found</span>
                                                    )}
                                                </td>
                                            </tr>
                                            <tr>
                                                <th className="text-secondary">Accessible Links</th>
                                                <td>
                                                    {result.AccessibleLinks > 0 ? result.AccessibleLinks 
                                                    : <span className="text-muted">No Accessible Links Found</span>}
                                                </td>
                                            </tr>
                                            <tr>
                                                <th className="text-secondary">Inaccessible Links</th>
                                                <td>
                                                    {result.InAccessibleLinks > 0 ? result.InAccessibleLinks 
                                                    : <span className="text-muted">No Inaccessible Links Found</span>}
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                    </div>
                )}
            </div>

            {/* Footer */}
            <footer className="bg-light text-center text-secondary py-3 mt-5 border-top small">
                <span>Web Page Analyzer &copy; {new Date().getFullYear()} &middot; <i className="bi bi-bootstrap-fill text-primary"></i> Sellaiya Suresh</span>
            </footer>
        </Fragment>
    );
}

export default AppClass;
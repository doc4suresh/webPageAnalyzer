import React, {Fragment} from 'react';

function AppClass() {
    return (
        <Fragment>
            <div className="container min-vh-100 d-flex align-items-center justify-content-center">
                <div className="w-100" style={{maxWidth: 420}}>
                    <div className="bg-white p-4 p-md-5 rounded-4 shadow text-center">
                        <h1 className="mb-3 text-primary fw-bold" style={{letterSpacing: 1}}>Web Page Analizer</h1>
                        <p className="mb-4 text-secondary">Enter the URL of the web page you want to analyze</p>
                        <input type="text" className="form-control form-control-lg mb-3" placeholder="Enter URL" />
                        <button className="btn btn-primary btn-lg w-100 fw-semibold" type="button">Analyze</button>
                    </div>
                </div>
            </div>
        </Fragment>
    )
}

export default AppClass;
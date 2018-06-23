import React  from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export const toastInProgress = (message) => {
    return toast(message, { autoClose: false });
};

export const toastSuccess = (message, toastID = null) => {
    if(toastID !== null) {
        toast.update(toastID, { render: message, type: toast.TYPE.SUCCESS, autoClose: 3000 });
    } else {
        toast.success(message);
    }
};

export const toastFail = (message, toastID = null) => {
    if(toastID !== null) {
        toast.update(toastID, { render: message, type: toast.TYPE.ERROR, autoClose: 3000 });
    } else {
        toast.error(message);
    }
};

export default () => {

    return (
        <ToastContainer position="top-center"
                        autoClose={5000}
                        hideProgressBar
                        newestOnTop={false}
                        closeOnClick
                        rtl={false}
                        pauseOnVisibilityChange
                        draggable
                        pauseOnHover />
    );
}

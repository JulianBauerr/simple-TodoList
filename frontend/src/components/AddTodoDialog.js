import {useRef, useState} from "react";
const AddTodoDialog = (props) => {
    const [isOpen, setIsOpen] = useState(false);
    const dialogRef = useRef(null);

    function handleSubmit(event) {
        event.preventDefault();
        const todoListId = parseInt(props.id);
        const name = document.getElementById("AddModal-" + props.id + "-name").value;
        const done = document.getElementById("AddModal-" + props.id + "-done").checked;
        const data = {todoListId, name, done};

        fetch('http://localhost:8000/todo/add', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
            })
            .catch((error) => {
                console.log('Error:', error);
            });

        setIsOpen(false)
        dialogRef.current.close();
        window.location.reload();
    }

    return (
        <dialog id={"AddModal-" + props.id} className="modal" ref={dialogRef} open={isOpen}>
            <div className="modal-box" style={{width: "min-content"}}>
                <h3 className="text-lg font-bold">Add a new ToDo</h3>
                <p className="py-4">Press ESC key to Cancel</p>
                <form onSubmit={handleSubmit}>
                    <p>Name:</p>
                    <input id={"AddModal-" + props.id + "-name"}/>
                    <p>Done:</p>
                    <input id={"AddModal-" + props.id + "-done"} type={"checkbox"} className={"checkbox"}/>
                    <div className="modal-action">
                        <button className="btn" type="submit">Add</button>
                    </div>
                </form>
            </div>
        </dialog>
    );
}

export default AddTodoDialog;
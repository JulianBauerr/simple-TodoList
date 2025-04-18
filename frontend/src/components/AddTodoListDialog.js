import {useRef, useState} from "react";
const AddTodoDialog = (props) => {
    const [isOpen, setIsOpen] = useState(false);
    const dialogRef = useRef(null);

    function handleSubmit(event) {
        event.preventDefault();
        const name = document.getElementById("AddTodoListDialog-" + props.id + "-name").value;
        const data = {name};
        console.log(data);
        fetch('http://localhost:8000/todoList/add', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        })
            .then(response => response.json())
            .then(data => {
                console.log(data)
                props.handler(props.id)
            })
            .catch(error => console.log(error))

        setIsOpen(false)
        dialogRef.current.close();
        window.location.reload();
    }

    return (
        <dialog id={"AddTodoListDialog-" + props.id} className="modal" ref={dialogRef} open={isOpen}>
            <div className="modal-box" style={{width: "min-content"}}>
                <h3 className="text-lg font-bold">Add a new ToDo</h3>
                <p className="py-4">Press ESC key to Cancel</p>
                <form onSubmit={handleSubmit}>
                    <p>Name:</p>
                    <input id={"AddTodoListDialog-" + props.id + "-name"}/>
                    <div className="modal-action">
                        <button className="btn" type="submit">Add</button>
                    </div>
                </form>
            </div>
        </dialog>
    );
}

export default AddTodoDialog;
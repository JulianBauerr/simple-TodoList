export const AddTodoListButton = (props) => {
    function AddTodoList() {
        const data = {}
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
    }

    return (
        <button id={"AddTodoListButton-" + props.id} className="btn btn-circle btn-outline"
                onClick={() => {
                    document.getElementById("AddTodoListDialog-"+props.id).showModal();
                }}
                tabIndex="-1"
                style={{
                    opacity: "0.4",
                    scale: "0.7",
                    fontSize: "20px",
                    fontWeight: "bold",
                }}
                onMouseOver={(e) => {
                    e.target.style.transform = "scale(1.2)";
                    e.target.style.transition = "0.5s";
                }}
                onMouseOut={(e) => {
                    e.target.style.transform = "scale(1)";
                    e.target.style.transition = "0.5s";
                }}
        >+
        </button>
    );
}
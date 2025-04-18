export const AddTodoButton = (props) =>{
    return (
        <div style={{
            display: "flex",
            justifyContent: "center",
        }}>
            <button id={"AddTodoButton-"+props.id} className="btn btn-circle btn-outline"
                    onClick={() => {
                        document.getElementById("AddModal-"+props.id).showModal();
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
        </div>
    );
}
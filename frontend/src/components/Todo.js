import React, {useState} from "react";

const Todo = (props) => {
    const [isCrossedOut, setIsCrossedOut] = useState(props.done);
    const [state, setState] = useState({
        id: props.id, name: props.title, done: props.done,
    });

    if (props.id < 0) {
        return null;
    }

    const toggleCheckbox = () => {
        setIsCrossedOut(!isCrossedOut);
        setState(prevState => ({
            ...prevState, done: !isCrossedOut
        }));

        const checkboxState = !isCrossedOut;
        const data = {checkboxState};
    };

    const width = `${props.length + 5}ch`;

    return (
        <div
            className={"shadow flex m-2 items-center"}
            style={{
                width: width,
                borderRadius: "50px",
                justifyContent: "space-between"
            }}>
            <p className={"p-2"} style={{textDecoration: isCrossedOut ? "line-through" : "none"}}>
                {props.title}
            </p>
            <input defaultChecked={state.done} type="checkbox" className="checkbox mr-2" onClick={toggleCheckbox}/>
        </div>
    );
}

export default Todo;
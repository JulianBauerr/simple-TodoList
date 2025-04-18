import {AddTodoButton} from "./AddTodoButton";
import Todo from "./Todo";
import AddTodoDialog from "./AddTodoDialog";
import {Reorder} from "framer-motion";
import {useEffect, useState} from "react";

export const TodoList = (props) => {
    useEffect(() => {
        setItems(props.todos);
    }, [props.todos]);

    const [items, setItems] = useState(props.todos);

    const handleReorder = (newOrder) => {
        const reorderedItems = newOrder.map(id => items.find(item => item.id === id));
        setItems(reorderedItems);

    };

    const todoLength = Math.max(...items.map(todo => todo.name.length));

    return (
        <Reorder.Group
            className="gap-1.5 flex flex-col m-8"
            axis="y"
            values={items.map(item => item.id)}
            onReorder={handleReorder}
        >
            {items.map((todo) => (
                <Reorder.Item
                    key={todo.id}
                    value={todo.id}>
                    <Todo
                        TodoListId={props.id}
                        title={todo.name}
                        id={todo.id}
                        done={todo.done}
                        length={todoLength}
                    />
                </Reorder.Item>
            ))}
            <AddTodoButton id={props.id}/>
            <AddTodoDialog id={props.id} handleLoad={props.handleLoad}/>
        </Reorder.Group>
    );
}

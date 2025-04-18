import Header from "./components/Header";
import {useEffect, useState} from "react";
import {TodoList} from "./components/TodoList";

const Home = (props) => {
    const [todoLists, setTodoLists] = useState([])

    function handleLoadAllTodoLists() {
        return fetch('http://localhost:8000/todoList/loadAll', {
            method: 'post'
        })
            .then(response => response.json())
            .then(data => {
                setTodoLists(data);
            })
            .catch((error) => {
                console.log('Error:', error);
            });
    }

    function handleLoadTodoList(id) {
        const data = {id: id};

        fetch('http://localhost:8000/todoList/load', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
            .then(response => response.json())
            .then(data => {
                updateTodoList({id: data.id, size: data.size, todos: data.todos})
            })
            .catch((error) => {
                console.log('Error:', error);
            });
    }

    function updateTodoList(TodoList) {
        setTodoLists(prevTodoLists => {
            const index = prevTodoLists.findIndex(list => list.id === TodoList.id);
            if (index !== -1) {
                const updatedTodoLists = [...prevTodoLists];
                updatedTodoLists[index] = TodoList;
                return updatedTodoLists;
            } else {
                return [...prevTodoLists, TodoList];
            }
        });
    }

    useEffect(() => {
        handleLoadAllTodoLists()
    }, []);

    return (
        <div data-theme="cupcake" style={{height: "100vh"}}>
            <div style={{
                justifyItems: "center",
                alignItems: "center",
                display: "flex",
                flexDirection: "column"
            }}>
                <Header handleLoad={handleLoadTodoList} handleLoadAll={handleLoadAllTodoLists} todoLists={todoLists}/>
            </div>
            <div className="todo-list-container" style={{
                display: "flex",
                flexWrap: "wrap",
                gap: "16px",
                padding: "16px"
            }}>
                {todoLists.map((todoList) => {
                    return <TodoList handleLoad={handleLoadTodoList} todos={todoList.todos}
                                     id={todoList.id} key={todoList.id}/>
                })}
            </div>
        </div>
    );
}

export default Home;
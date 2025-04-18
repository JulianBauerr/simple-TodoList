import {Link} from "react-router-dom";
import { AddTodoListButton } from './AddTodoListButton';
import AddTodoListDialog from "./AddTodoListDialog";

const Header = ({ handleLoadAll, todoLists }) => {
    return (
        <div className="navbar bg-base-100 shadow" style={{
            width: "80vw",
            borderRadius: "25px",
            height: "8vh"
        }}>
            <div className="flex-1">
                <Link to={"/"} className="btn btn-ghost text-xl">Home</Link>
            </div>
            <div className="flex-row">
                <ul className="menu menu-horizontal px-1">
                    {todoLists.map((todoList) => {
                        return <li key={todoList.id}><Link to={"/List_id=" + todoList.id}>{todoList.name}</Link></li>
                    })}
                </ul>
                <AddTodoListButton id={0} handler={handleLoadAll}/>
                <AddTodoListDialog id={0}/>
            </div>
        </div>
    );
}

export default Header;
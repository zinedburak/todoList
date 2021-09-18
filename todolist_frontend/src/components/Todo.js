import React from "react";

const Todo = ( {todo} ) => {
    return (
        <div className="todo">
            <li id="todo-item" className="todo-item" data-testid="todo" > {todo["content"]} </li>
        </div>
    );
};

export default Todo
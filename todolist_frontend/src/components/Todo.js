import React from "react";

const Todo = ( {todo} ) => {
    return (
        <div className="todo">
            <li className="todo-item" data-testid="todo" > {todo["text"]} </li>
        </div>
    );
};

export default Todo
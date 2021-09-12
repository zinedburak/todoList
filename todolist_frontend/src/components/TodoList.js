import React from "react";
import Todo from "./Todo";

const TodoList = ( { todos } ) => {
    return (
        <div className="todo-container">
            <ul className="todo-list" data-testid="todoList" >
                {todos.map((todo) => (
                    <Todo
                    key = {todo["id"]}
                    todo = {todo}
                    
                    />
                ))}
            </ul>

        </div>
    );
};

export default TodoList
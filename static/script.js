const checkIcon = `<svg class="icon icon--check" xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px"><path d="m424-296 282-282-56-56-226 226-114-114-56 56 170 170Zm56 216q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z"/></svg>`
const circleIcon = `<svg class="icon icon--circle" xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z"/></svg>`
const deleteIcon = `<svg class="icon icon--delete" xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px"><path d="M280-120q-33 0-56.5-23.5T200-200v-520h-40v-80h200v-40h240v40h200v80h-40v520q0 33-23.5 56.5T680-120H280Zm400-600H280v520h400v-520ZM360-280h80v-360h-80v360Zm160 0h80v-360h-80v360ZM280-720v520-520Z"/></svg>`

const todoItems = document.getElementById('todo-items')
const newTodoInput = document.getElementById('new-todo-input')
const newTodoButton = document.getElementById('new-todo-button')

const connection = new WebSocket(`ws://${document.location.host}/ws`)

let todos = []
let todoItemElements = []

connection.onopen = () => connection.send(`{"action":"get_todos"}`)
connection.onmessage = event => renderTodos(event.data)

newTodoButton.addEventListener('click', () => {
    message = JSON.stringify({
        action: 'new',
        description: newTodoInput.value
    })
    connection.send(message)

    newTodoInput.value = null
})

function renderTodos(data) {
    todoItemElements.forEach(element => {
        element.remove()
    })

    todoItemElements = []

    if (!data) return

    todos = JSON.parse(data)

    todos.forEach(item => {
        const todoItem = document.createElement('li')

        todoItem.classList.add('item')
        todoItem.classList.add(item.Complete ?? 'item--complete')

        const toggleButton = document.createElement('button')
        toggleButton.classList.add('item__button')
        toggleButton.classList.add('item__toggle-button')
        toggleButton.classList.add(item.Complete ? 'item__toggle-button--complete' : 'item__toggle-button--incomplete')
        toggleButton.innerHTML = item.Complete ? checkIcon : circleIcon
        toggleButton.onclick = () => {
            message = JSON.stringify({
                action: 'toggle',
                id: item.ID,
            })
            connection.send(message)
        }
        
        const description = document.createElement('p')
        description.classList.add('item__description')
        description.innerText = item.Description
        
        const deleteButton = document.createElement('button')
        deleteButton.classList.add('item__button')
        deleteButton.classList.add('item__delete-button')
        deleteButton.innerHTML = deleteIcon
        deleteButton.onclick = () => {
            message = JSON.stringify({
                action: 'delete',
                id: item.ID,
            })
            connection.send(message)
        }

        todoItem.appendChild(toggleButton)
        todoItem.appendChild(description)
        todoItem.appendChild(deleteButton)

        todoItemElements.push(todoItem)
    })

    todoItemElements.forEach(element => {
        todoItems.appendChild(element)
    })
}

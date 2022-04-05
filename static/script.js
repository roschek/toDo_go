const notes = document.querySelector('.todos')
const sendingBtn = document.querySelector('.form__button')

const getData=() => {
  fetch('/note')
  .then(res => res.json())
  .then(data=> {    
    if(data) {
      console.log(data);
    data.forEach(element => {
      notes.insertAdjacentHTML('beforeend',
      `
      <div class="todo__row">
        <p class="todo__item">${element.title}</p>
        <p class="todo__item">${element.topic}</p>
      </div>
      `)
    });
  }
  })
}

getData()

// sendingBtn.addEventListener('click',(evt)=>{
//   evt.preventDefault()
//   const myForm = document.forms.todo
//   const title = myForm.elements.title.value
//   const topic = myForm.elements.topic.value  
//   let note = new FormData()
//   note.append('title',title)
//   note.append('topic', topic)
//   fetch('/note',{
//     method: "POST",
//     body: note
//   })
//   .then(getData())
// })

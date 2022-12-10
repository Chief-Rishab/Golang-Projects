import React,{Component} from "react";
import axios, { formToJSON } from 'axios'
import {Card, Header, Form, Input, Icon, Button, CardContent, CardHeader} from "semantic-ui-react";

let endpoint ="http://localhost:8000";

class ToDoList extends Component{
    constructor(props){
        super(props);

        this.state={
            task:"",
            items:[],
        }
    }

    componenetDidMount(){
        this.getTask();
    }

    onChange = (event) => {
        this.setState({
            [event.target.name] : event.target.value
        });
    };

    onSubmit = () =>{
        console.log("create task called");
        let {task} = this.state;
        console.log(task);
        if (task){
            axios.post(endpoint+ "/api/article",
            {
                "url":task,
            },
            {
                headers:
                {
                "Content-Type":"application/json",
                },
            }
    ).then((res)=>{
        this.getTask();
        this.setState({
            task:"",
        });
        console.log(res);
    });
    }
    };

    getTask = ()=>{
        console.log("Get task called");
        axios.get(endpoint+"/api/articles").then((res)=>{
            console.log(res);
            if(res.data){
                this.setState({
                    items: res.data.map((item)=>{
                        console.log(res);
                        let color = "yellow";
                        let style={
                            wordWrap:"break-word",
                        };
                      
                    if(item.read){
                        color="green";
                        style["textDecorationLine"]="underline";
                    }    
                    
                    return(

                        <Card key={item._id} color={color} fluid className="rough">
                            <CardContent>
                                <CardHeader textAlign="left">
                                    <div style={style}>{item.title}</div>
                                </CardHeader>
                                <div >{item.description}</div>
                                <div>
                                <img src={item.imageurl} style={{ width: '200px', }} />
                                </div>
                                <Button href={item.url} target="_blank">Read More</Button>
                                <Card.Meta textAlign="left">
                               
                                </Card.Meta>
                                <Card.Meta textAlign="right">        
                                <Icon
                                     name="check circle"
                                     color="green"
                                     onClick={() => this.updateTask(item._id)}
                                 />
                                <span style={{ paddingRight: 10 }}>Done</span>
                                    <Icon 
                                    name="undo"
                                    color="blue"
                                    onClick={()=> this.updateUnread(item._id)}
                                    />
                                    <span style={{paddingRight:10}}> Undo</span>
                                    <Icon 
                                    name="delete"
                                    color="red"
                                    onClick={()=> this.deleteTask(item._id)}
                                    />
                                    <span style={{paddingRight:10}}>Delete</span>
                                </Card.Meta>

                            </CardContent>

                        </Card>
                    );
                    }),
                });
            }else{
                this.setState({
                    items:[]
                });
            }
        });
    };

    updateTask = (id) =>{
        console.log("update task called");
        axios.put(endpoint+"/api/article/"+id,{
            headers:{
                "Content-Type":"application/json",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    };

    updateUnread = (id) =>{
        console.log("update task as unread called");
        axios.put(endpoint+"/api/articleUnread/"+id,{
            headers:{
                "Content-Type":"application/json",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    }

    deleteTask = (id) =>{
        console.log("delete task called");
        axios.delete(endpoint+"/api/delArticle/"+id,{
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    };

    deleteAllTask = () =>{
        console.log("delete all task called");
        axios.delete(endpoint+"/api/deletearticles",{
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    };

    render(){
        return(
            <div>
                <div className="row">
                    <Header className="header" as="h2" color="black">
                        ReadNote
                    </Header>
                    
                </div>
                <div className="row">
                        <Form onSubmit={this.onSubmit}>
                        <Input
                        type="text"
                        name="task"
                        onChange={this.onChange}
                        value={this.state.task}
                        fluid
                        placeholder="Enter any website URL here..."
                        />
                        {<Button>Create Article</Button>}
                        </Form>
                </div>
                <div>
                    <Button onClick={this.deleteAllTask}>Clear All</Button>
                </div>
                <div className="row">
                    <Card.Group>{this.state.items}</Card.Group>
                </div>
            </div>
        );
    }
}

export default ToDoList;
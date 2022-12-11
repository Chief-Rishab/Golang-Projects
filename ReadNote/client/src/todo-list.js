import React,{Component} from "react";
import axios from 'axios'
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
        let {task} = this.state;
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
        //console.log(res);
    });
    }
    };

    getTask = ()=>{
        axios.get(endpoint+"/api/articles").then((res)=>{
            //console.log(res);
            if(res.data){
                this.setState({
                    items: res.data.map((item)=>{
                    //    console.log(res);
                        let color = "yellow";
                        let style={
                            wordWrap:"break-word",
                        };
                      
                    if(item.read){
                        color="green";
                        style["textDecorationLine"]="underline";
                        style["textDecorationColor"]="blue";
                    }    
                    
                    return(

                        <Card key={item._id} color={color} fluid className="rough" style={{ width: '25rem', borderBottom: 'solid', backgroundColor:{color} }} >
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
                                     onClick={() => this.updateRead(item._id)}
                                 />
                                <span style={{ paddingRight: 10 }}>Mark as Read</span>
                                    <Icon 
                                    name="undo"
                                    color="yellow"
                                    onClick={()=> this.updateUnread(item._id)}
                                    />
                                    <span style={{paddingRight:10}}> UnRead</span>
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

    updateRead = (id) =>{
        axios.put(endpoint+"/api/article/"+id,{
            headers:{
                "Content-Type":"application/json",
            },
        }).then((res)=>{
           // console.log(res);
            this.getTask();
        });
    };

    updateUnread = (id) =>{
        axios.put(endpoint+"/api/articleUnread/"+id,{
            headers:{
                "Content-Type":"application/json",
            },
        }).then((res)=>{
           // console.log(res);
            this.getTask();
        });
    }

    deleteTask = (id) =>{
        axios.delete(endpoint+"/api/delArticle/"+id,{
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
          //  console.log(res);
            this.getTask();
        });
    };

    deleteAllTask = () =>{
        axios.delete(endpoint+"/api/deletearticles",{
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
         //   console.log(res);
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
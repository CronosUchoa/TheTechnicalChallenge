import './Home.css';
import Display from './Display.jsx';
import Condicional from './Condicional.jsx';
import React , { Component } from 'react';
import axios from 'axios';


axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';


class Todos extends Component{

  //criando estado 
    state = {
      address :[],
      confirmations:[]
    }
  
    handleChange = event => {
      console.log(this.setState({ address: event.target.value }));
    }

    handleSubmit = event => {
      event.preventDefault();


    axios.get(`http://localhost:8000/balance/?address=${this.state.address}`)
      .then(res => {
        // console.log(res);
        // console.log(res.data);
        this.setState({confirmations: res.data})
      })
  
      
 
   }
  
      render(){

              const {confirmations} = this.state;
            
            
            return(

                    <div>
                      <form onSubmit={this.handleSubmit} >
                      <h1 className="tituloTexto">Bitcoin</h1>
                      <input className="imputTexto" type="text" name="address" onChange={this.handleChange}/>
                      <button type="submit" className="botao" >Enviar</button>
                      </form>
                        
                          {confirmations.map(confCoin=>(
                            
                          <Condicional numberConfir={confCoin.confirmations} key={confCoin.txid} valor={confCoin.value} nome={" "}></Condicional>
                            // <div key={confCoin.txid}>
                            //   <Display valor={confCoin.confirmations} nome={"Saldo acumulado confirmado"}></Display>
                            // </div>


                          ))}
                  
                    </div>
           
                  ) 

      }
      
    
}

export default Todos;
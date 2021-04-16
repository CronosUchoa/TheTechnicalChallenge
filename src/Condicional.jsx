import React, { useState } from 'react'
import Display from './Display.jsx';

let cond = props => {


    return(
        <div>     
            {
                props.numberConfir < 2  ?
                <div key={props.txid}> 
                              <Display valor={props.valor} nome={"Unconfirmed Balance"}></Display>
                </div>
                    :
                    <Display valor={props.valor} nome={"Confirmed Balance"}></Display>
            }
        </div>

    )
}

export default cond;
import React from "react";
import NavbarUser from "../components/navbar";
import { Container, Row, Col, Modal } from "react-bootstrap";
import { DataTransaction } from "../datadummy/datatransaction";
import Icon from "../assets/TrashIcon.png";
import convertRupiah from "rupiah-format";
import { useState } from "react";
import { API } from "../config/api";
import { useNavigate } from "react-router-dom";
import { useContext } from "react";
import { Usercontext } from "../context/usercontext";
import { useEffect } from "react";
import { useMutation } from "react-query";
import ModalCart from "../components/Trans";

export default function Cart() {
  const [state, dispatch] = useContext(Usercontext);
  const [carts, setCarts] = useState([])
  // modal
  const [showTrans, setShowTrans] = useState(false);
  // const handleShow = () => setShowTrans(true);
  const handleClose = () => setShowTrans(false);
  let navigate = useNavigate();

  // cart
  // let { data: cart, refetch } = useQuery("cartsCache", async () => {
  //   const response = await API.get("/carts-id");
  //   return response.data.data;
  // });
  const getCart = async () => {
    try {
      const response = await API.get("/carts-id");
      setCarts(response.data.data);
    } catch (error) {
      console.log(error);
    }
  };

  console.log(carts);

  useEffect(() => {
    getCart();
  }, []);


  let resultTotal = carts?.reduce((a, b) => {
    return a + b.subtotal;
  }, 0);
  console.log(resultTotal)

  const form = {
    status: "pending",
    total: resultTotal,
  };

  const handleSubmit = useMutation(async (e) => {
    const config = {
      headers: {
        "Content-type": "application/json",
      },
    };
    // Insert transaction data
    const body = JSON.stringify(form);

    const response = await API.patch("/transaction", body, config);
    console.log(response.data.data.token);

    const token = response.data.data.token;

    window.snap.pay(token, {
      onSuccess: function (result) {
        /* You may add your own implementation here */
        console.log(result);
        navigate("/profile");
      },
      onPending: function (result) {
        /* You may add your own implementation here */
        console.log(result);
        navigate("/profile");
      },
      onError: function (result) {
        /* You may add your own implementation here */
        console.log(result);
      },
      onClose: function () {
        /* You may add your own implementation here */
        alert("you closed the popup without finishing the payment");
      },
    });
  });
  

  useEffect(() => {
    //change this to the script source you want to load, for example this is snap.js sandbox env
    const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
    //change this according to your client-key
    const myMidtransClientKey = "Client key here ...";

    let scriptTag = document.createElement("script");
    scriptTag.src = midtransScriptUrl;
    // optional if you want to set script attribute
    // for example snap.js have data-client-key attribute
    scriptTag.setAttribute("data-client-key", myMidtransClientKey);

    document.body.appendChild(scriptTag);
    return () => {
      document.body.removeChild(scriptTag);
    };
  }, []);

  return (
    <div>
      <NavbarUser />
      <div>
        <Container>
          <div className="ms-5 mt-3">
            <h1 style={{ color: "#BD0707" }}>My Cart</h1>
            <Row>
              <p style={{ color: "#BD0707" }}>Review Your Order</p>
              <Col md={8}>
                <hr />

                {carts.map((item, index) => {
                  return (
                    <div className="mb-2">
                      <Row>
                        <Col key={index} md={2}>
                          <img
                            src={item?.product?.image}
                            alt=""
                            style={{ width: "100%" }}
                          />
                        </Col>
                        <Col md={10}>
                          <Col className="d-flex justify-content-between">
                            <p>
                              <strong style={{ color: "#BD0707" }}>
                                {item?.product?.title}
                              </strong>
                            </p>
                            <p>{convertRupiah.convert(item?.product?.price)}</p>
                          </Col>
                          <Col className="d-flex justify-content-between">
                            <span >Topping :</span>
                            
                          
                            {item?.Topping?.map((topping, idx) => (
                              
                              <span className="d-inline" key={idx}>{topping.title}{","}</span>
                            ))}
                            
                            
                            <img
                              src={Icon}
                              alt=""
                              style={{ width: "20px", height: "20px" }}
                            />
                          </Col>
                        </Col>
                      </Row>
                    </div>
                  );
                })}
                <hr />
              </Col>
              <Col md={4}>
                <hr />
                {carts.map((item, index) => {
                  return(
                    <>
                <Col className="d-flex justify-content-between">
                  <p>SubTotal</p>
                  <p>{convertRupiah.convert(item.subtotal)}</p>
                </Col>
                <Col className="d-flex justify-content-between">
                  <p>Qty</p>
                  {/* <p>{qty}</p> */}
                </Col>
                </>
                  )
                })}
                <hr />
                
                

                <Col className="d-flex justify-content-between">
                  <p>Total</p>
                  <p>{convertRupiah.convert(resultTotal)}</p>
                </Col>
                <button
                  type="button"
                  className="pt-2 pb-2"
                  style={{
                    width: "100%",
                    color: "white",
                    backgroundColor: "red",
                    borderColor: "red",
                    borderRadius: "5px",
                  }}
                   onClick={(e) => handleSubmit.mutate(e)}
                >
                  Pay
                </button>
              </Col>
            </Row>
          </div>
        </Container>
      </div>
      <div
        className="modal fade"
        id="thanks-for-order"
        tabindex="-1"
        role="dialog"
        aria-labelledby="exampleModalCenterTitle"
        aria-hidden="true"
      ></div>
      <ModalCart showTrans={showTrans} close={handleClose} />
    </div>
  );
}

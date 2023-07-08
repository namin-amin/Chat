import {MessageTypes} from "./messages";
import {client} from "../utils/client";

type onmessageFn = (Id: MessageEvent) => void;

export class SSE extends EventSource {
    public sseId = "";
    //maintains the map of subscriptions of each event
    private Subscriptions: Map<string, onmessageFn[]>;

    /**
     *@param url represents the url to be connected for SSE server endpoint
     * @param userID is the id of the currently connecting use
     */
    constructor(url: string, userID: string) {
        super(url);
        this.sseId = userID;
        this.Subscriptions = new Map<string, onmessageFn[]>();
        this.Subscribe(MessageTypes.Open.toString(), (ev) => {
            this.verify(ev, userID)
        });
        console.log("subscribes");

    }

    private verify(e: MessageEvent, userId: string) {
        console.log("run the verify")
        const initData = JSON.parse(e.data);
        console.log(initData);

        client.post("/chats/sse/verify", {
            userId: userId,
            sseId: initData.data
        }).then(() => {
            console.log("verified sse");

        }).catch((err) => {
            console.log(err
            );
        });
    }

    /**
     * Subscribe to different events
     */
    Subscribe(event: string, callback: onmessageFn) {
        if (event === "") {
            console.error("event string is empty");
            return;
        }
        if (callback === undefined) {
            console.error("callback function cannot be empty");
            return;
        }
        if (this.Subscriptions.has(event)) {
            return;
            //this.Subscriptions.get(event)?.push(callback); //TODO fix the repeated calling
        } else {
            console.log(event + " event added")
            this.Subscriptions.set(event, [callback]);
            this.addEventListener(event, (e) => {
                this.Subscriptions.get(event)?.forEach(callBackFunc => {
                    console.log(e.data);
                    callBackFunc(e);
                });
            });
        }
    }
}
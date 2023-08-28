import { getScaleFreeNetwork } from "../js/util.js";

var network;
var container;
var exportArea;

function init() {
    container = document.getElementById("mynetwork");
    exportArea = document.getElementById("input_output");
}

function addConnections(elem, index) {
    // need to replace this with a tree of the network, then get child direct children of the element
    elem.connections = network.getConnectedNodes(index);
}

// Clear graph view
function destroyNetwork() {
    network.destroy();
}

function clearOutputArea() {
    exportArea.value = "";
}

function getNodeData(data) {
    var networkNodes = [];

    data.forEach(function (elem, index, array) {
        networkNodes.push({
            id: elem.id,
            label: elem.id,
            x: elem.x,
            y: elem.y,
        });
    });

    return new vis.DataSet(networkNodes);
}

function getNodeById(data, id) {
    for (var n = 0; n < data.length; n++) {
        if (data[n].id == id) {
            // Double equals since id can be numeric or string
            return data[n];
        }
    }

    throw "Can not find id '" + id + "' in data";
}

function getEdgeData(data) {
    var networkEdges = [];

    data.forEach(function (node) {
        // add the connection
        node.connections.forEach(function (connId, cIndex, conns) {
            networkEdges.push({ from: node.id, to: connId });
            let cNode = getNodeById(data, connId);

            var elementConnections = cNode.connections;

            // Remove the connection from the other node to prevent duplicate connections
            var duplicateIndex = elementConnections.findIndex(function (
                connection
            ) {
                return connection == node.id; // Double equals since id can be numeric or string
            });

            if (duplicateIndex != -1) {
                elementConnections.splice(duplicateIndex, 1);
            }
        });
    });

    return new vis.DataSet(networkEdges);
}

init();






var clusterIndex = 0;
var clusters = [];
var lastClusterZoomLevel = 0;
var clusterFactor = 0.9;

// create an array with nodes
const nodes = [
    {
        id: 1,
        label: "User 1",
        group: "users",
    },
    {
        id: 2,
        label: "User 2",
        group: "users",
    },
    {
        id: 3,
        label: "Usergroup 1",
        group: "usergroups",
    },
    {
        id: 4,
        label: "Usergroup 2",
        group: "usergroups",
    },
    {
        id: 5,
        label: "Organisation 1",
        shape: "icon",
        icon: {
            face: "'Ionicons'",
            code: "\uf276",
            size: 50,
            color: "#f0a30a",
        },
    },
    {
        id: 6,
        label: "User 3",
        group: "users",
    },
    {
        id: 7,
        label: "User 4",
        group: "users",
    },
];

// create an array with edges
const edges = [
    { from: 1, to: 3 },
    { from: 1, to: 4 },
    { from: 2, to: 4 },
    { from: 3, to: 5 },
    { from: 4, to: 5 },
    { from: 6, to: 7 },
    { from: 6, to: 5 },
    { from: 5, to: 2 },
    { from: 7, to: 3 },
];

// create a network

var data = {
    nodes: nodes,
    edges: edges,
};
var options = {
    layout: { randomSeed: 8 },
    physics: { adaptiveTimestep: false },
    manipulation: { enabled: true }, // Turn this off for production
    groups: {
        usergroups: {
            shape: "icon",
            icon: {
                face: "'Ionicons'",
                code: "\uf47c",
                size: 50,
                color: "#57169a",
            },
        },
        users: {
            shape: "icon",
            icon: {
                face: "'Ionicons'",
                code: "\uf47e",
                size: 50,
                color: "#aa00ff",
            },
        },
    }
};
var network = new vis.Network(container, data, options);
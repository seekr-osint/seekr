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
var nodes = [
    { id: 1, label: "Node 1" },
    { id: 2, label: "Node 2" },
    { id: 3, label: "Node 3" },
    { id: 4, label: "Node 4" },
    { id: 5, label: "Node 5" },
    { id: 6, label: "Node 6" },
    { id: 7, label: "Node 7" },
    { id: 8, label: "Node 8" },
    { id: 9, label: "Node 9" },
    { id: 10, label: "Node 10" },
];

// create an array with edges
var edges = [
    { from: 1, to: 2 },
    { from: 1, to: 3 },
    { from: 10, to: 4 },
    { from: 2, to: 5 },
    { from: 6, to: 2 },
    { from: 7, to: 5 },
    { from: 8, to: 6 },
    { from: 9, to: 7 },
    { from: 10, to: 9 },
];

// create a network

var data = {
    nodes: nodes,
    edges: edges,
};
var options = {
    layout: { randomSeed: 8 },
    physics: { adaptiveTimestep: false },
    manipulation: { enabled: true }
};
var network = new vis.Network(container, data, options);

// set the first initial zoom level
network.once("initRedraw", function () {
    if (lastClusterZoomLevel === 0) {
        lastClusterZoomLevel = network.getScale();
    }
});

// we use the zoom event for our clustering
network.on("zoom", function (params) {
    if (params.direction == "-") {
        if (params.scale < lastClusterZoomLevel * clusterFactor) {
            makeClusters(params.scale);
            lastClusterZoomLevel = params.scale;
        }
    } else {
        openClusters(params.scale);
    }
});

// if we click on a node, we want to open it up!
network.on("selectNode", function (params) {
    if (params.nodes.length == 1) {
        if (network.isCluster(params.nodes[0]) == true) {
            network.openCluster(params.nodes[0]);
            network.stabilize();
        }
    }
});

// make the clusters
function makeClusters(scale) {
    var clusterOptionsByData = {
        processProperties: function (clusterOptions, childNodes) {
            clusterIndex = clusterIndex + 1;
            var childrenCount = 0;
            for (var i = 0; i < childNodes.length; i++) {
                childrenCount += childNodes[i].childrenCount || 1;
            }
            clusterOptions.childrenCount = childrenCount;
            clusterOptions.label = "# " + childrenCount + "";
            clusterOptions.font = { size: childrenCount * 5 + 30 };
            clusterOptions.id = "cluster:" + clusterIndex;
            clusters.push({ id: "cluster:" + clusterIndex, scale: scale });
            return clusterOptions;
        },
        clusterNodeProperties: {
            borderWidth: 3,
            shape: "database",
            font: { size: 30 },
        },
    };
    network.clusterOutliers(clusterOptionsByData);

    network.setOptions({ physics: { stabilization: { fit: false } } });
    network.stabilize();
}

// open them back up!
function openClusters(scale) {
    var newClusters = [];
    for (var i = 0; i < clusters.length; i++) {
        if (clusters[i].scale < scale) {
            network.openCluster(clusters[i].id);
            lastClusterZoomLevel = scale;
        } else {
            newClusters.push(clusters[i]);
        }
    }
    clusters = newClusters;

    network.stabilize();
}